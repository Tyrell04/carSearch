<script lang="ts">
	import { profileStore } from '$lib/stores/profile';
	import { 
		Button,
		Alert, 
		AlertTitle, 
		AlertDescription,
		Stack,
		DonationButton
	} from '$lib/components';

	let profile = $derived($profileStore);

	function dismissBanner() {
		profileStore.dismissDonationBanner();
	}

	function openDonationLink() {
		window.open('https://buymeacoffee.com/carsearch', '_blank', 'noopener,noreferrer');
	}
</script>

{#if !profile.donationBannerDismissed}
	<Alert class="border-yellow-200 bg-yellow-50 text-yellow-800 mb-6">
		<Stack direction="row" justify="between" align="center" class="w-full">
			<Stack direction="row" gap="sm" align="center" class="flex-1">
				<svg class="h-5 w-5 text-yellow-600" viewBox="0 0 24 24" fill="currentColor">
					<path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 15h2v2h-2v-2zm0-8h2v6h-2V9z"/>
				</svg>
				<Stack gap="xs" class="flex-1">
					<AlertTitle class="text-sm font-semibold text-yellow-800">
						🚀 Unterstütze die Serverkosten!
					</AlertTitle>
					<AlertDescription class="text-sm text-yellow-700">
						Diese App läuft auf einem Server, der monatliche Kosten verursacht. Mit einer kleinen Spende hilfst du dabei, den Service am Laufen zu halten und weitere Features zu entwickeln. Vielen Dank für deine Unterstützung! ☕
					</AlertDescription>
				</Stack>
			</Stack>
			<Stack direction="row" gap="sm" class="ml-4">
				<DonationButton size="sm" variant="outline" class="bg-white hover:bg-yellow-100 text-yellow-700 border-yellow-300" />
				<Button 
					variant="ghost" 
					size="sm" 
					onclick={dismissBanner}
					class="text-yellow-600 hover:text-yellow-800 hover:bg-yellow-100 h-8 w-8 p-0"
					aria-label="Banner schließen"
				>
					<svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
						<path d="M18 6L6 18M6 6l12 12"/>
					</svg>
				</Button>
			</Stack>
		</Stack>
	</Alert>
{/if}
