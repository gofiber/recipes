<script>
	import { CodeBlock } from '@skeletonlabs/skeleton';
	import { gofiber_code, robustRouting, serverStaticFiles } from './lib/store/code';
</script>

<svelte:head>
	<title>Fiber & Sveltekit Embed App</title>
</svelte:head>
<section class="flex flex-col">
	<div class="flex flex-col justify-center p-4">
		<h1 class="h1">An Express-inspired web framework written in Go.</h1>
		<blockquote class="blockquote my-2">
			<p>
				Fiber is a Go web framework built on top of Fasthttp, the <strong>fastest</strong> HTTP
				engine for Go. It's designed to <strong>ease</strong> things up for fast development with
				<strong>zero memory allocation</strong> and <strong>performance</strong> in mind.
			</p>
		</blockquote>
		<div class="flex flex-col lg:flex-row">
			<!---Code Block-->
			<div class="w-full h-96 flex flex-col">
				<div class="flex-grow">
					<CodeBlock lineNumbers={true} rounded="rounded-t-lg" language="go" code={$gofiber_code} />
				</div>
			</div>
			<!---Code Block End-->
			<div class="flex justify-start items-center">
				<span class="font-extrabold text-5xl select-none rotate-90 lg:rotate-0 m-1">&rarr;</span>
			</div>
			<!---Browser Mockup-->
			<div class="w-full max-h-full flex flex-col">
				<div
					class="w-full h-11 rounded-t-lg bg-gray-200 flex justify-start items-center space-x-1.5 px-3"
				>
					<span class="relative w-3 h-3 border-2 rounded-full border-red-400" />
					<span class="relative w-3 h-3 border-2 rounded-full border-yellow-400" />
					<span class="relative w-3 h-3 border-2 rounded-full border-green-400" />
					<div class="border flex-grow">
						<span class="text-base">http://localhost:3000</span>
					</div>
				</div>
				<div class="card p-4 h-auto flex-grow">Hello, World!</div>
			</div>
			<!---Browser Mockup End-->
		</div>
	</div>
</section>
<section class="w-full mb-8 mt-8">&nbsp;</section>
<section class="flex h-full flex-col lg:flex-row bg-primary-200">
	<div class="w-full inline-flex flex-col justify-center items-center">
		<h3 class="h3">Robust Routing</h3>
		<p class="p-3">
			Setting up routes for your application has never been so easy! The Express-like route definitions are easy to understand and work with.
		</p>
	</div>
	<div class="w-full flex flex-col">
		<CodeBlock lineNumbers={true} rounded="rounded-t-lg" language="go" code={$robustRouting} />
	</div>
</section>
<section class="w-full mb-8 mt-8">&nbsp;</section>
<section class="flex h-full flex-col lg:flex-row bg-primary-200">
	<div class="w-full flex flex-col">
		<CodeBlock lineNumbers={true} rounded="rounded-t-lg" language="go" code={$serverStaticFiles} />
	</div>
	<div class="w-full inline-flex flex-col justify-center items-center">
		<h3 class="h3">Serve Static Files</h3>
		<p class="p-3">
			Serve your static HTML, CSS, and JavaScript files with ease by defining static routes. You can also serve the contents of multiple directories on the same route!
		</p>
	</div>
</section>
<section class="w-full mb-8 mt-8">&nbsp;</section>