import './style.scss';
import {Form, InputType} from '../../components/Form';
import {useState} from "preact/hooks";
import {carQuery} from "../../lib/api";
import {Table} from "../../components/Table";

export function Home() {

	const [data, setData] = useState([]);

	const handleSubmit = async (event) => {
		event.preventDefault();
		const formData = new FormData(event.target);
		const hsn = formData.get('hsn');
		const tsn = formData.get('tsn');
		const response = await carQuery({hsn, tsn});
		console.log(response.name);

		setData(data.concat(response));
	};

	return (
		<main class="container">
			<hgroup>
				<h1>carSearch</h1>
				<h2>Typenschlüsselnummern und Herstellerschlüsselnummern Suche</h2>
			</hgroup>
		<Form onSubmit={handleSubmit}>
			<Form.Input type={InputType.TEXT} name="hsn" placeholder="Typenschlüsselnummer" />
			<Form.Input type={InputType.TEXT} name="tsn" placeholder="Herstellerschlüsselnummer" />
			<Form.Submit>Suchen</Form.Submit>
		</Form>
			<Table>
				<Table.Header columns={["Autotype", "Hersteller", "Typenschlüsselnummer", "Herstellerschlüsselnummer"]} />
				{data.map((car) => (
					<Table.Row column={[car.name, car.manufacturer_name, car.hsn, car.tsn]} />
				))
				}
			</Table>
		</main>
	);
}