<script lang="ts">
	import { cn } from "$lib/utils";

	interface Props {
		children: import("svelte").Snippet;
		title?: string;
		subtitle?: string;
		actions?: import("svelte").Snippet;
		class?: string;
	}

	let {
		children,
		title,
		subtitle,
		actions,
		class: className = "",
		...restProps
	}: Props = $props();
</script>

<header class={cn("mb-8", className)} {...restProps}>
	<div class="flex items-center justify-between">
		<div class="flex-1"></div>
		<div class="flex-1 text-center">
			{#if title}
				<h1 class="text-3xl font-bold tracking-tight">{title}</h1>
			{/if}
			{#if subtitle}
				<p class="text-sm text-muted-foreground">{subtitle}</p>
			{/if}
		</div>
		<div class="flex-1 flex justify-end">
			{#if actions}
				{@render actions()}
			{/if}
		</div>
	</div>
	{#if children}
		{@render children()}
	{/if}
</header>
