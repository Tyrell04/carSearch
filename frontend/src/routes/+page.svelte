<script lang="ts">
	import type { Car, Producer } from '$lib/types';
	import { api } from '$lib/api';
	import { profileStore } from '$lib/stores/profile';
	import { 
		Button, 
		Input, 
		Card, 
		CardHeader, 
		CardTitle, 
		CardContent,
		Alert, 
		AlertTitle, 
		AlertDescription,
		Container,
		Stack,
		Grid,
		Header,
		ThemeSwitcher,
		ProfileModal,
		SavedCarsList,
		// DonationBanner, // Uncomment this line to enable donation banner
		DonationButton
	} from '$lib/components';

	let hsn = $state('');
	let tsn = $state('');
	let results: Car[] = $state([]);
	let producer: Producer | null = $state(null);
	let loading = $state(false);
	let error = $state('');
	let showProfileModal = $state(false);

	let profile = $derived($profileStore);

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
			const data = await api.searchCars(hsn.trim(), tsn.trim() || undefined);
			
			if (data.car) {
				results = [data.car];
				producer = null;
			} else if (data.producer) {
				producer = data.producer;
				results = [];
			} else if (data.cars) {
				results = data.cars;
				producer = null;
			} else {
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

	function saveCar(car: Car) {
		profileStore.saveCar(car);
	}

	function isCarSaved(car: Car): boolean {
		return profile.savedCars.some(
			(savedCar) => savedCar.HSN === car.HSN && savedCar.TSN === car.TSN
		);
	}

	function openProfileModal() {
		showProfileModal = true;
	}

	function closeProfileModal() {
		showProfileModal = false;
	}
</script>

<svelte:head>
	<title>Autosuche</title>
</svelte:head>

<Container size="lg" class="py-10">
	<Stack gap="lg">
		<!-- Donation Banner -->
		<!-- <DonationBanner /> --> <!-- Uncomment this line to enable donation banner -->

		<!-- Header with Actions -->
		<Header 
			title="Autosuche"
			subtitle="Suche nach Autos anhand der HSN und optional der TSN."
		>
			{#snippet actions()}
				<Stack direction="row" gap="sm">
					<ThemeSwitcher />
					<DonationButton size="sm" variant="outline" />
					<Button 
						variant="outline" 
						size="sm" 
						onclick={() => window.location.href = '/saved'}
					>
						<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
							<path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.29 1.51 4.04 3 5.5l7 7Z" />
						</svg>
						Gespeicherte ({profile.savedCars.length})
					</Button>
					<Button 
						variant="outline" 
						size="sm" 
						onclick={openProfileModal}
					>
						<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
							<path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
							<circle cx="12" cy="7" r="4" />
						</svg>
						Profil
					</Button>
				</Stack>
			{/snippet}
		</Header>

		<!-- Search Form -->
		<Card>
			<form onsubmit={(e) => { e.preventDefault(); searchCars(); }}>
				<CardHeader>
					<CardTitle class="text-lg">Suchkriterien</CardTitle>
				</CardHeader>
				<CardContent>
					<Stack gap="md">
						<Stack gap="sm">
							<label for="hsn" class="text-sm font-medium">HSN (erforderlich)</label>
							<Input id="hsn" bind:value={hsn} placeholder="HSN eingeben" required disabled={loading} />
						</Stack>
						<Stack gap="sm">
							<label for="tsn" class="text-sm font-medium">TSN (optional)</label>
							<Input id="tsn" bind:value={tsn} placeholder="TSN eingeben" disabled={loading} />
						</Stack>
						<Stack direction="row" gap="sm" class="sm:flex-row flex-col">
							<Button type="submit" disabled={loading || !hsn.trim()}>
								{loading ? 'Suche...' : 'Autos suchen'}
							</Button>
							<Button type="button" variant="secondary" onclick={clearSearch} disabled={loading}>
								Löschen
							</Button>
						</Stack>
					</Stack>
				</CardContent>
			</form>
		</Card>

		<!-- Error Display -->
		{#if error}
			<Alert variant="destructive">
				<AlertTitle>Fehler</AlertTitle>
				<AlertDescription>{error}</AlertDescription>
			</Alert>
		{/if}

		<!-- Search Results -->
		{#if results.length}
			<Stack gap="md">
				<h2 class="text-xl font-semibold">
					Suchergebnisse ({results.length} {results.length === 1 ? 'Auto' : 'Autos'} gefunden)
				</h2>
				<Grid cols="auto" gap="md">
					{#each results as car, index}
						<Card>
							<CardHeader class="pb-2">
								<Stack direction="row" justify="between" align="center">
									<CardTitle class="text-base">Auto #{index + 1}</CardTitle>
									<Button
										variant={isCarSaved(car) ? "default" : "outline"}
										size="sm"
										onclick={() => saveCar(car)}
										disabled={isCarSaved(car)}
										class="h-7 px-2 text-xs"
									>
										{isCarSaved(car) ? 'Gespeichert' : 'Speichern'}
									</Button>
								</Stack>
							</CardHeader>
							<CardContent class="pt-0 text-sm">
								<Stack gap="xs">
									<div><span class="font-medium inline-block w-28">HSN:</span>{car.HSN}</div>
									<div><span class="font-medium inline-block w-28">TSN:</span>{car.TSN}</div>
									<div><span class="font-medium inline-block w-28">Fahrzeug:</span>{car.Name}</div>
									<div><span class="font-medium inline-block w-28">Hersteller:</span>{car.Producer?.Name || 'Unbekannt'}</div>
								</Stack>
							</CardContent>
						</Card>
					{/each}
				</Grid>
			</Stack>
		{:else if producer}
			<Card>
				<CardHeader class="pb-2">
					<CardTitle class="text-lg">Herstellerinformationen</CardTitle>
				</CardHeader>
				<CardContent class="pt-0 text-sm">
					<Stack gap="xs">
						<div><span class="font-medium inline-block w-32">HSN:</span>{producer.HSN}</div>
						<div><span class="font-medium inline-block w-32">Hersteller:</span>{producer.Name}</div>
					</Stack>
				</CardContent>
			</Card>
		{:else if !loading && (hsn || tsn)}
			<p class="text-center text-sm italic text-muted-foreground">
				Keine Autos für die angegebenen Kriterien gefunden.
			</p>
		{/if}

		<!-- Additional Components -->
		<SavedCarsList />
	</Stack>
</Container>

<!-- Profile Modal -->
<ProfileModal open={showProfileModal} onClose={closeProfileModal} />

<style>
	/* Clean, semantic styling using shadcn/ui components for better maintainability */
</style>
