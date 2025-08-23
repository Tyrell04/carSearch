<!-- SavedCarsList.svelte -->
<script lang="ts">
	import { profileStore } from '$lib/stores/profile';
	import { 
		Card, 
		CardHeader, 
		CardTitle, 
		CardContent,
		Button,
		Stack,
		Grid
	} from '$lib/components';
	import { exportCarsToCSV } from '$lib/utils/export';
	import type { Car } from '$lib/types';
	
	let profile = $derived($profileStore);
	
	function removeCar(car: Car) {
		profileStore.removeSavedCar(car);
	}
	
	function clearAllCars() {
		if (confirm('Möchten Sie wirklich alle gespeicherten Autos löschen?')) {
			profileStore.clearSavedCars();
		}
	}
	
	function exportToCSV() {
		exportCarsToCSV(profile.savedCars, `saved-cars-${new Date().toISOString().split('T')[0]}.csv`);
	}
</script>

{#if profile.showSavedCars && profile.savedCars.length > 0}
	<Card>
		<CardHeader class="pb-3">
			<Stack direction="row" justify="between" align="center">
				<CardTitle class="text-lg">
					Gespeicherte Autos ({profile.savedCars.length})
				</CardTitle>
				<Stack direction="row" gap="sm">
					<Button
						variant="outline"
						size="sm"
						onclick={exportToCSV}
						class="h-7 px-2 text-xs"
					>
						CSV Export
					</Button>
					<Button
						variant="destructive"
						size="sm"
						onclick={clearAllCars}
						class="h-7 px-2 text-xs"
					>
						Alle löschen
					</Button>
				</Stack>
			</Stack>
		</CardHeader>
		
		<CardContent class="pt-0">
			<Grid cols="auto" gap="md">
				{#each profile.savedCars as car, index}
					<Card class="border-muted/60">
						<CardHeader class="pb-2">
							<Stack direction="row" justify="between" align="center">
								<CardTitle class="text-base">Gespeichert #{index + 1}</CardTitle>
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
		</CardContent>
	</Card>
{/if}
