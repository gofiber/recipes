<script>
	import { onMount } from 'svelte';

	let dark = false;

	onMount(() => {
		dark =
			localStorage.getItem('theme') === 'dark' ||
			(localStorage.getItem('theme') === null &&
				window.matchMedia('(prefers-color-scheme: dark)').matches);
		apply();
	});

	function apply() {
		document.documentElement.classList.toggle('dark', dark);
	}

	function toggle() {
		dark = !dark;
		localStorage.setItem('theme', dark ? 'dark' : 'light');
		apply();
	}
</script>

<button
	type="button"
	on:click={toggle}
	aria-label="Toggle dark mode"
	aria-pressed={dark}
	class="rounded-full border border-primary-400 px-3 py-1 text-lg leading-none transition-colors hover:bg-primary-50 dark:hover:bg-primary-900"
>
	{dark ? '☀️' : '🌙'}
</button>
