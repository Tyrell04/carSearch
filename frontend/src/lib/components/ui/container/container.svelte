<script lang="ts">
	import { cn } from "$lib/utils";

	interface Props {
		children: import("svelte").Snippet;
		size?: "sm" | "md" | "lg" | "xl" | "full";
		class?: string;
	}

	let {
		children,
		size = "lg",
		class: className = "",
		...restProps
	}: Props = $props();

	const sizeClasses = {
		sm: "max-w-sm",
		md: "max-w-md", 
		lg: "max-w-3xl",
		xl: "max-w-6xl",
		full: "max-w-full"
	};
</script>

<div class={cn("container mx-auto px-4", sizeClasses[size], className)} {...restProps}>
	{@render children()}
</div>
