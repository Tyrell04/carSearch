<script lang="ts">
	import { cn } from "$lib/utils";

	interface Props {
		children: import("svelte").Snippet;
		direction?: "row" | "col";
		gap?: "none" | "xs" | "sm" | "md" | "lg" | "xl";
		align?: "start" | "center" | "end" | "stretch";
		justify?: "start" | "center" | "end" | "between" | "around" | "evenly";
		class?: string;
	}

	let {
		children,
		direction = "col",
		gap = "md",
		align = "stretch",
		justify = "start",
		class: className = "",
		...restProps
	}: Props = $props();

	const directionClasses = {
		row: "flex-row",
		col: "flex-col"
	};

	const gapClasses = {
		none: "gap-0",
		xs: "gap-1",
		sm: "gap-2", 
		md: "gap-4",
		lg: "gap-6",
		xl: "gap-8"
	};

	const alignClasses = {
		start: "items-start",
		center: "items-center",
		end: "items-end",
		stretch: "items-stretch"
	};

	const justifyClasses = {
		start: "justify-start",
		center: "justify-center",
		end: "justify-end",
		between: "justify-between",
		around: "justify-around",
		evenly: "justify-evenly"
	};
</script>

<div class={cn(
	"flex",
	directionClasses[direction],
	gapClasses[gap],
	alignClasses[align],
	justifyClasses[justify],
	className
)} {...restProps}>
	{@render children()}
</div>
