<script lang="ts">
	import { mode, setMode } from "mode-watcher";
	import { Button } from "../button";
	
	interface Props {
		size?: "default" | "sm" | "lg" | "icon";
		variant?: "default" | "destructive" | "outline" | "secondary" | "ghost" | "link";
		class?: string;
	}
	
	let {
		size = "icon",
		variant = "outline",
		class: className = "",
		...restProps
	}: Props = $props();
</script>

<Button
	{variant}
	{size}
	class={className}
	onclick={() => setMode($mode === "dark" ? "light" : "dark")}
	{...restProps}
>
	<!-- Sun Icon -->
	<svg class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" 
		fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
		<circle cx="12" cy="12" r="5"/>
		<path d="M12 1v2m0 18v2M4.22 4.22l1.42 1.42m12.72 12.72l1.42 1.42M1 12h2m18 0h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/>
	</svg>
	<!-- Moon Icon -->
	<svg class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" 
		fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
		<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
	</svg>
	<span class="sr-only">Theme wechseln</span>
</Button>
