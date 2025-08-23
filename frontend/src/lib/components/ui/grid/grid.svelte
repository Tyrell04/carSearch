<script lang="ts">
	import { cn } from "$lib/utils";

	interface Props {
		children: import("svelte").Snippet;
		cols?: 1 | 2 | 3 | 4 | 6 | 12 | "auto";
		gap?: "none" | "xs" | "sm" | "md" | "lg" | "xl";
		responsive?: boolean;
		class?: string;
	}

	let {
		children,
		cols = "auto",
		gap = "md",
		responsive = true,
		class: className = "",
		...restProps
	}: Props = $props();

	const colClasses = {
		1: "grid-cols-1",
		2: "grid-cols-2",
		3: "grid-cols-3",
		4: "grid-cols-4",
		6: "grid-cols-6",
		12: "grid-cols-12",
		auto: "grid-cols-1 md:grid-cols-2"
	};

	const gapClasses = {
		none: "gap-0",
		xs: "gap-1",
		sm: "gap-2",
		md: "gap-4",
		lg: "gap-6",
		xl: "gap-8"
	};
</script>

<div class={cn(
	"grid",
	colClasses[cols],
	gapClasses[gap],
	className
)} {...restProps}>
	{@render children()}
</div>
