<script>
	import hljs from 'highlight.js';

	export let code = '';
	export let language = 'plaintext';
	export let lineNumbers = false;
	export let rounded = 'rounded-lg';

	// Trailing newlines would render as a stray empty gutter row.
	$: trimmed = code.replace(/\n+$/, '');
	$: highlighted = hljs.highlight(trimmed, { language }).value;
	$: lineCount = trimmed.split('\n').length;
</script>

<div class="flex h-full overflow-hidden bg-gray-900 text-sm {rounded}">
	{#if lineNumbers}
		<div
			aria-hidden="true"
			class="shrink-0 select-none border-r border-gray-700 px-3 py-4 text-right font-mono leading-6 text-gray-500"
		>
			{#each Array(lineCount) as _, i}
				<div>{i + 1}</div>
			{/each}
		</div>
	{/if}
	<pre class="flex-grow overflow-auto px-4 py-4 leading-6"><code class="language-{language}"
			>{@html highlighted}</code
		></pre>
</div>
