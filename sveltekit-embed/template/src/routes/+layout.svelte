<script lang="ts">
	import 'highlight.js/styles/github-dark.css';
	// Most of your app wide CSS should be put in this file
	import '../app.css';

	import LightSwitch from '$lib/LightSwitch.svelte';
</script>

<div class="flex h-full flex-col overflow-hidden">
	<header
		class="grid shrink-0 grid-cols-3 items-center border-b border-gray-200 px-4 py-2 dark:border-gray-700"
	>
		<img src="/fiber-logo.svg" alt="fiber-logo" class="w-24 justify-self-start" />
		<div></div>
		<div class="justify-self-end">
			<LightSwitch />
		</div>
	</header>

	<div class="flex-grow overflow-y-auto">
		<div class="p-1">
			<slot />
		</div>

		<footer class="flex flex-col items-center justify-center bg-primary-50 p-4">
			<p class="mb-4 text-3xl font-extrabold text-primary-400">...and much more!</p>
			<a
				href="https://docs.gofiber.io/"
				class="inline-flex items-center rounded-lg border border-primary-500 px-4 py-2 transition-colors hover:bg-primary-100"
				data-sveltekit-preload-data="hover"
			>
				<span class="text-primary-900">Explore Fiber</span>
				<span class="m-1 rotate-90 select-none text-5xl font-extrabold text-primary-900 lg:rotate-0"
					>&rarr;</span
				>
			</a>
		</footer>
	</div>
</div>
