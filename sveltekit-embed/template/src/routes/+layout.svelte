<script lang="ts">
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-rocket.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	// Most of your app wide CSS should be put in this file

	import 'highlight.js/styles/github-dark.css';

	import '../app.postcss';

	import hljs from 'highlight.js';
	import { LightSwitch, AppBar, AppShell, storeHighlightJs } from '@skeletonlabs/skeleton';
	storeHighlightJs.set(hljs);
</script>

<AppShell>
	<svelte:fragment slot="header">
		<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
			<svelte:fragment slot="lead">
				<img src="/fiber-logo.svg" alt="fiber-logo" class="w-24" />
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<LightSwitch />
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<div class="p1">
		<slot />
	</div>
	<svelte:fragment slot="pageFooter">
		<div class="p-4 flex flex-col justify-center items-center bg-primary-50">
			<p class="font-extrabold text-3xl text-primary-400 mb-4">...and much more!</p>
			<a
				href="https://docs.gofiber.io/"
				class="btn variant-outline-primary"
				data-sveltekit-preload-data="hover"
			>
				<span class="text-primary-900">Explore Fiber</span>
				<span class="text-primary-900 font-extrabold text-5xl select-none rotate-90 lg:rotate-0 m-1"
					>&rarr;</span
				>
			</a>
		</div>
	</svelte:fragment>
</AppShell>
