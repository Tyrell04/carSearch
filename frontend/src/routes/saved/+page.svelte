<script lang="ts">
	import { profileStore } from '$lib/stores/profile';
	import { 
		Card, 
		CardHeader, 
		CardTitle, 
		CardContent,
		Button,
		Alert, 
		AlertTitle, 
		AlertDescription,
		Container,
		Stack,
		Grid,
		Header,
		ThemeSwitcher
	} from '$lib/components';
	import { exportCarsToCSV } from '$lib/utils/export';
	import type { Car } from '$lib/types';
	
	let profile = $derived($profileStore);
	
	function removeCar(car: Car) {
		if (confirm('Möchten Sie dieses Auto wirklich entfernen?')) {
			profileStore.removeSavedCar(car);
		}
	}
	
	function clearAllCars() {
		if (confirm('Möchten Sie wirklich alle gespeicherten Autos löschen?')) {
			profileStore.clearSavedCars();
		}
	}
	
	function exportToCSV() {
		exportCarsToCSV(profile.savedCars, `saved-cars-${new Date().toISOString().split('T')[0]}.csv`);
	}
	
	function goBack() {
		history.back();
	}
</script>

<svelte:head>
	<title>Gespeicherte Autos - Autosuche</title>
</svelte:head>

<Container size="xl" class="py-10">
	<Stack gap="lg">
		<!-- Header with Actions -->
		<Header
			title="Gespeicherte Autos"
			subtitle={`Verwalten Sie Ihre gespeicherten Fahrzeuge (${profile.savedCars.length} ${profile.savedCars.length === 1 ? 'Auto' : 'Autos'})`}
		>
			{#snippet actions()}
				<Stack direction="row" gap="sm">
					<ThemeSwitcher />
					<Button variant="outline" size="sm" onclick={goBack}>
						<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
							<path d="M19 12H5M12 19l-7-7 7-7" />
						</svg>
						Zurück
					</Button>
				</Stack>
			{/snippet}
		</Header>
		
		<!-- Action Buttons -->
		{#if profile.savedCars.length > 0}
			<Stack direction="row" gap="sm" class="flex-wrap">
				<Button variant="outline" size="sm" onclick={exportToCSV}>
					<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
						<path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
						<polyline points="7,10 12,15 17,10" />
						<line x1="12" y1="15" x2="12" y2="3" />
					</svg>
					Als CSV exportieren
				</Button>
				<Button variant="destructive" size="sm" onclick={clearAllCars}>
					<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
						<polyline points="3,6 5,6 21,6" />
						<path d="m19,6v14a2,2 0 0,1 -2,2H7a2,2 0 0,1 -2,-2V6m3,0V4a2,2 0 0,1 2,-2h4a2,2 0 0,1 2,2v2" />
					</svg>
					Alle löschen
				</Button>
			</Stack>
		{/if}

		<!-- Content -->
		{#if profile.savedCars.length === 0}
			<Alert>
				<svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
					<circle cx="12" cy="12" r="10" />
					<path d="m9,9 6,6" />
					<path d="m15,9 -6,6" />
				</svg>
				<AlertTitle>Keine gespeicherten Autos</AlertTitle>
				<AlertDescription>
					Sie haben noch keine Autos gespeichert. Gehen Sie zur Suche und speichern Sie interessante Fahrzeuge.
				</AlertDescription>
			</Alert>
		{:else}
			<Stack gap="lg">
				<!-- Cars Grid -->
				<Grid cols={3} gap="md">
					{#each profile.savedCars as car, index}
						<Card class="hover:shadow-md transition-shadow">
							<CardHeader class="pb-3">
								<Stack direction="row" justify="between" align="center">
									<CardTitle class="text-base">
										<Stack direction="row" gap="sm" align="center">
											<svg class="h-4 w-4 text-blue-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
												<path d="M3 12h18m-9-9v18" />
											</svg>
											Auto #{index + 1}
										</Stack>
									</CardTitle>
									<Button
										variant="ghost"
										size="sm"
										onclick={() => removeCar(car)}
										class="h-6 w-6 p-0 text-muted-foreground hover:text-destructive"
										aria-label="Auto entfernen"
									>
										<svg class="h-3 w-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
											<path d="M18 6L6 18M6 6l12 12" />
										</svg>
									</Button>
								</Stack>
							</CardHeader>
							<CardContent class="pt-0">
								<Stack gap="md">
									<Stack gap="sm" class="text-sm">
										<div class="flex">
											<span class="font-medium inline-block w-20 text-muted-foreground">HSN:</span>
											<span class="font-mono">{car.HSN}</span>
										</div>
										<div class="flex">
											<span class="font-medium inline-block w-20 text-muted-foreground">TSN:</span>
											<span class="font-mono">{car.TSN}</span>
										</div>
									</Stack>
									
									<div class="border-t pt-3">
										<Stack gap="sm">
											<div>
												<span class="text-xs text-muted-foreground uppercase tracking-wide">Fahrzeug</span>
												<p class="font-medium">{car.Name}</p>
											</div>
											<div>
												<span class="text-xs text-muted-foreground uppercase tracking-wide">Hersteller</span>
												<p class="text-sm">{car.Producer?.Name || 'Unbekannt'}</p>
											</div>
										</Stack>
									</div>
									
									{#if car.CreatedAt}
										<div class="text-xs text-muted-foreground pt-2 border-t">
											Gespeichert: {new Date(car.CreatedAt).toLocaleDateString('de-DE')}
										</div>
									{/if}
								</Stack>
							</CardContent>
						</Card>
					{/each}
				</Grid>
				
				<!-- Summary -->
				<div class="text-center">
					<p class="text-sm text-muted-foreground">
						Insgesamt {profile.savedCars.length} {profile.savedCars.length === 1 ? 'Auto' : 'Autos'} gespeichert
					</p>
				</div>
			</Stack>
		{/if}
	</Stack>
</Container>

<style>
	/* Clean, semantic styling using shadcn/ui components for better maintainability */
</style>
