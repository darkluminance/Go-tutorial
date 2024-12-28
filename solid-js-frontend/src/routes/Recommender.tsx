import { createEffect, createSignal, onMount, Show } from "solid-js";
import { getAuthenticatedData, postAuthenticatedData } from "../utils";
import DashboardBase from "../components/DashboardBase";
import DynamicTable from "../components/DynamicTable";
import { Recommenders } from "../models/Recommender";
import toast from "solid-toast";
import ItemInputBase from "../components/ItemInputBase";
import RecommenderForm from "../components/Inputs/RecommenderForm";
import BaseAllEntries from "../components/BaseAllEntries";

export default function Recommender() {
	return <DashboardBase component={RecommenderComponent} />;
}

function RecommenderComponent() {
	const [recommenders, setRecommenders] = createSignal<Recommenders[]>();
	const [showItemInputter, setShowItemInputter] = createSignal(false);
	const [editData, setEditData] = createSignal();

	const fetchRecommenders = async () => {
		setRecommenders();

		const recommenders = (await getAuthenticatedData(
			"/recommenders/user/get"
		)) as Recommenders[];

		if (recommenders) {
			setRecommenders(recommenders);
		}

		return recommenders;
	};
	const handleDeleteRecommender = (id: string) => {
		const res = postAuthenticatedData("/recommenders/delete", {
			id: id,
		});

		toast
			.promise(res, {
				loading: "Deleting data...",
				success: (val) => <span>Deleted data</span>,
				error: <span>Could not delete data</span>,
			})
			.then(() => {
				fetchRecommenders();
			});
	};

	const handleCompleteSubmit = () => {
		setShowItemInputter(false);
		fetchRecommenders();
	};

	onMount(() => {
		toast.promise(fetchRecommenders(), {
			loading: "Fetching data...",
			success: (val) => <span>Fetched data</span>,
			error: <span>Could not fetch data</span>,
		});
	});

	return (
		<>
			<BaseAllEntries
				title="Recommenders"
				data={recommenders()}
				handleDelete={handleDeleteRecommender}
				formComponent={RecommenderForm}
				fetchData={fetchRecommenders}
			/>
		</>
	);
}
