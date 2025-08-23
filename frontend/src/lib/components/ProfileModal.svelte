<!-- ProfileModal.svelte -->
<script lang="ts">
	import { profileStore } from '$lib/stores/profile';
	import { 
		Button,
		Card, 
		CardHeader, 
		CardTitle, 
		CardContent,
		Stack
	} from '$lib/components';
	import type { Car } from '$lib/types';
	
	interface Props {
		open: boolean;
		onClose: () => void;
	}
	
	let { open, onClose }: Props = $props();
	
	let profile = $derived($profileStore);
	
	function toggleShowSavedCars() {
		profileStore.toggleShowSavedCars();
	}
	
	// Close modal when clicking outside
	function handleBackdropClick(e: MouseEvent) {
		if (e.target === e.currentTarget) {
			onClose();
		}
	}
	
	// Handle keyboard events on backdrop
	function handleBackdropKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter' && e.target === e.currentTarget) {
			onClose();
		}
	}
	
	// Close modal on Escape key
	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape') {
			onClose();
		}
	}
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open}
	<!-- Modal backdrop -->
	<div 
		class="fixed inset-0 z-50 bg-black/50 backdrop-blur-sm"
		onclick={handleBackdropClick}
		onkeydown={handleBackdropKeydown}
		role="dialog"
		aria-modal="true"
		aria-labelledby="profile-modal-title"
		tabindex="-1"
	>
		<!-- Modal content -->
		<div class="fixed left-1/2 top-1/2 z-50 w-full max-w-lg -translate-x-1/2 -translate-y-1/2 transform p-4">
			<Card class="w-full shadow-lg">
				<CardHeader class="pb-4">
					<Stack direction="row" justify="between" align="center">
						<CardTitle id="profile-modal-title" class="text-lg">Profil Einstellungen</CardTitle>
						<Button 
							variant="ghost" 
							size="sm" 
							onclick={onClose}
							class="h-8 w-8 p-0"
							aria-label="Schließen"
						>
							<svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
								<path d="M18 6L6 18M6 6l12 12" />
							</svg>
						</Button>
					</Stack>
				</CardHeader>
				
				<CardContent>
					<Stack gap="lg">
						<!-- Display Mode Toggle -->
						<Stack gap="sm">
							<h3 class="text-sm font-medium">Ansicht-Modus</h3>
							<Stack direction="row" gap="sm">
								<Button
									variant={profile.showSavedCars ? "default" : "outline"}
									size="sm"
									onclick={toggleShowSavedCars}
									class="flex-1"
								>
									Liste anzeigen
								</Button>
								<Button
									variant={!profile.showSavedCars ? "default" : "outline"}
									size="sm"
									onclick={toggleShowSavedCars}
									class="flex-1"
								>
									Liste verstecken
								</Button>
							</Stack>
							<p class="text-xs text-muted-foreground">
								{profile.showSavedCars 
									? 'Gespeicherte Autos werden in einer Liste angezeigt' 
									: 'Gespeicherte Autos sind versteckt'
								}
							</p>
						</Stack>
						
						<!-- Saved Cars Section -->
						<Stack gap="sm">
							<Stack direction="row" justify="between" align="center">
								<h3 class="text-sm font-medium">
									Gespeicherte Autos ({profile.savedCars.length})
								</h3>
								<Button
									variant="outline"
									size="sm"
									onclick={() => { onClose(); window.location.href = '/saved'; }}
									class="h-7 px-2 text-xs"
								>
									Alle anzeigen
								</Button>
							</Stack>
							
							<p class="text-xs text-muted-foreground">
								{profile.savedCars.length === 0 
									? 'Keine gespeicherten Autos' 
									: `${profile.savedCars.length} ${profile.savedCars.length === 1 ? 'Auto' : 'Autos'} gespeichert. Klicken Sie auf "Alle anzeigen" für die vollständige Verwaltung.`
								}
							</p>
						</Stack>
					</Stack>
				</CardContent>
			</Card>
		</div>
	</div>
{/if}
