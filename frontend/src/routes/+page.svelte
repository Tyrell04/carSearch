<script lang="ts">
	import type { Car, Producer, ApiResponse } from '$lib/types';

	let hsn = '';
	let tsn = '';
	let results: Car[] = [];
	let producer: Producer | null = null;
	let loading = false;
	let error = '';

	// Assuming the API runs on the same host but different port or path
	// if in dev mode, you might need to adjust this URL in dev localhost:3000 and prod car.marc-schulz.online
	const isDev = import.meta.env.DEV;
	const API_BASE = isDev ? 'http://localhost:3000/api' : 'https://car.marc-schulz.online/api';

	async function searchCars() {
		if (!hsn.trim()) {
			error = 'HSN ist erforderlich';
			return;
		}

		loading = true;
		error = '';
		results = [];
		producer = null;

		try {
			// Build query parameters
			const params = new URLSearchParams();
			params.set('hsn', hsn.trim());

			if (tsn.trim()) {
				params.set('tsn', tsn.trim());
			}

			const url = `${API_BASE}/cars/search?${params.toString()}`;

			const response = await fetch(url);
			const data: ApiResponse = await response.json();

			if (!response.ok) {
				throw new Error(data.message || 'Suche fehlgeschlagen');
			}

			// Handle single car response (HSN + TSN search)
			if (data.car) {
				results = [data.car];
				producer = null;
			}
			// Handle producer response (HSN only search)
			else if (data.producer) {
				producer = data.producer;
				results = [];
			}
			// Handle multiple cars response (fallback)
			else if (data.cars) {
				results = data.cars;
				producer = null;
			}
			else {
				results = [];
				producer = null;
			}

		} catch (err) {
			error = err instanceof Error ? err.message : 'Ein Fehler ist aufgetreten';
		} finally {
			loading = false;
		}
	}

	function clearSearch() {
		hsn = '';
		tsn = '';
		results = [];
		producer = null;
		error = '';
	}
</script>

<svelte:head>
	<title>Autosuche</title>
</svelte:head>

<div class="container">
	<h1>Autosuche</h1>
	<p>Suche nach Autos anhand der HSN (erforderlich) und optional der TSN für spezifischere Ergebnisse.</p>

	<form on:submit|preventDefault={searchCars}>
		<div class="form-group">
			<label for="hsn">HSN (erforderlich):</label>
			<input
				id="hsn"
				type="text"
				bind:value={hsn}
				placeholder="HSN eingeben"
				required
				disabled={loading}
			/>
		</div>

		<div class="form-group">
			<label for="tsn">TSN (optional):</label>
			<input
				id="tsn"
				type="text"
				bind:value={tsn}
				placeholder="TSN für spezifisches Auto eingeben"
				disabled={loading}
			/>
		</div>

		<div class="buttons">
			<button type="submit" disabled={loading || !hsn.trim()}>
				{loading ? 'Suche...' : 'Autos suchen'}
			</button>
			<button type="button" on:click={clearSearch} disabled={loading}>
				Löschen
			</button>
		</div>
	</form>

	{#if error}
		<div class="error">
			<strong>Fehler:</strong> {error}
		</div>
	{/if}

	{#if results.length > 0}
		<div class="results">
			<h2>Suchergebnisse ({results.length} {results.length === 1 ? 'Auto' : 'Autos'} gefunden)</h2>

			{#each results as car, index}
				<div class="car-card">
					<h3>Auto #{index + 1}</h3>
					<div class="car-details">
						<div><strong>HSN:</strong> {car.HSN}</div>
						<div><strong>TSN:</strong> {car.TSN}</div>
						<div><strong>Fahrzeugname:</strong> {car.Name}</div>
						<div><strong>Hersteller:</strong> {car.Producer?.Name || 'Unbekannt'}</div>
					</div>
				</div>
			{/each}
		</div>
	{:else if producer}
		<div class="producer-result">
			<h2>Herstellerinformationen</h2>
			<div class="producer-details">
				<div><strong>HSN:</strong> {producer.HSN}</div>
				<div><strong>Herstellername:</strong> {producer.Name}</div>
			</div>
		</div>
	{:else if !loading && (hsn || tsn)}
		<div class="no-results">
			<p>Keine Autos für die angegebenen Kriterien gefunden.</p>
		</div>
	{/if}
</div>

<style>
	.container {
		max-width: 800px;
		margin: 0 auto;
		padding: 2rem;
		font-family: Arial, sans-serif;
	}

	h1 {
		color: #333;
		text-align: center;
		margin-bottom: 1rem;
	}

	p {
		text-align: center;
		color: #666;
		margin-bottom: 2rem;
	}

	form {
		background: #f5f5f5;
		padding: 2rem;
		border-radius: 8px;
		margin-bottom: 2rem;
	}

	.form-group {
		margin-bottom: 1rem;
	}

	label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: bold;
		color: #333;
	}

	input {
		width: 100%;
		padding: 0.75rem;
		border: 1px solid #ddd;
		border-radius: 4px;
		font-size: 1rem;
		box-sizing: border-box;
	}

	input:focus {
		outline: none;
		border-color: #007bff;
		box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
	}

	input:disabled {
		background-color: #e9ecef;
		opacity: 0.6;
	}

	.buttons {
		display: flex;
		gap: 1rem;
		margin-top: 1.5rem;
	}

	button {
		padding: 0.75rem 1.5rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		font-size: 1rem;
		transition: background-color 0.2s;
	}

	button[type="submit"] {
		background-color: #007bff;
		color: white;
	}

	button[type="submit"]:hover:not(:disabled) {
		background-color: #0056b3;
	}

	button[type="button"] {
		background-color: #6c757d;
		color: white;
	}

	button[type="button"]:hover:not(:disabled) {
		background-color: #545b62;
	}

	button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}

	.error {
		background-color: #f8d7da;
		color: #721c24;
		padding: 1rem;
		border-radius: 4px;
		margin-bottom: 1rem;
		border: 1px solid #f5c6cb;
	}

	.results {
		margin-top: 2rem;
	}

	.results h2 {
		color: #333;
		margin-bottom: 1rem;
	}

	.car-card {
		background: white;
		border: 1px solid #ddd;
		border-radius: 8px;
		padding: 1.5rem;
		margin-bottom: 1rem;
		box-shadow: 0 2px 4px rgba(0,0,0,0.1);
	}

	.car-card h3 {
		margin-top: 0;
		margin-bottom: 1rem;
		color: #007bff;
	}

	.car-details div {
		margin-bottom: 0.5rem;
		line-height: 1.4;
	}

	.car-details strong {
		color: #333;
		min-width: 120px;
		display: inline-block;
	}

	.no-results {
		text-align: center;
		padding: 2rem;
		color: #666;
		font-style: italic;
	}

	.producer-result {
		background: #e9f7ef;
		padding: 1.5rem;
		border-radius: 8px;
		margin-top: 2rem;
	}

	.producer-result h2 {
		color: #333;
		margin-bottom: 1rem;
	}

	.producer-details div {
		margin-bottom: 0.5rem;
		line-height: 1.4;
	}

	.producer-details strong {
		color: #333;
		min-width: 120px;
		display: inline-block;
	}

	@media (max-width: 600px) {
		.container {
			padding: 1rem;
		}

		.buttons {
			flex-direction: column;
		}
	}
</style>
