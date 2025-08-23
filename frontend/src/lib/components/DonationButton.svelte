<script lang="ts">
	import { Button } from '$lib/components';

	interface Props {
		size?: 'sm' | 'default' | 'lg';
		variant?: 'default' | 'outline' | 'secondary';
		class?: string;
	}

	let { 
		size = 'default', 
		variant = 'default', 
		class: className = '' 
	}: Props = $props();

	function openDonationLink() {
		window.open('https://buymeacoffee.com/carsearch', '_blank', 'noopener,noreferrer');
	}
</script>

<Button 
	{variant} 
	{size} 
	onclick={openDonationLink}
	class="bg-yellow-500 hover:bg-yellow-600 text-white border-yellow-500 hover:border-yellow-600 {className}"
>
	<svg class="h-4 w-4 mr-2" viewBox="0 0 24 24" fill="currentColor">
		<path d="M20 3H4v10c0 2.21 1.79 4 4 4h6c2.21 0 4-1.79 4-4v-3h2c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2zm-2 10c0 1.1-.9 2-2 2H8c-1.1 0-2-.9-2-2V5h12v8zM6 19h12v2H6z"/>
	</svg>
	<span class="text-sm font-medium">☕ Buy Me a Coffee</span>
</Button>
