import './style.scss';
import {Form, InputType} from '../../components/Form';
import {useState} from "preact/hooks";
import {carQuery} from "../../lib/api";

export function Home() {

	const [data, setData] = useState(null);

	const handleSubmit = async (event) => {
		event.preventDefault();
		const formData = new FormData(event.target);
		const hsn = formData.get('hsn');
		const tsn = formData.get('tsn');
		const response = await carQuery({hsn, tsn});
		setData(response);
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
		{data && (
			<table>
				<thead>
					<tr>
						<th>Typenschlüsselnummer</th>
						<th>Herstellerschlüsselnummer</th>
						<th>Modell</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>{data.hsn}</td>
						<td>{data.tsn}</td>
						<td>{data.name}</td>
					</tr>
				</tbody>
			</table>
		)}
		</main>
	);
}