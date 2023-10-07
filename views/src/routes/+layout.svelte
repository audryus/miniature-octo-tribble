<script>
	import '../app.postcss';
	import empresaStore from '../Store.js';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { Toast, getToastStore } from '@skeletonlabs/skeleton';
	import { initializeStores } from '@skeletonlabs/skeleton';

	initializeStores();

	/** @type {import('./$types').PageData} */
	export let data;
	
	let empresas = 0

	empresaStore.update((currentData) => {
		return data?.empresas;
	});
	empresaStore.subscribe((data) => {
		empresas = data.length;
	});

</script>

<Toast />
<!-- App Shell -->
<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="sidebarLeft">
	<!-- Insert the list: -->
	<nav class="list-nav">
		<ul>
			<li><a href="/">Home</a></li>
			<li><a href="/lista-ativos">Lista de ativos</a></li>
			<li class="relative">
					<a href="/empresas">
						Inconsistencias
					</a>
					{#if empresas > 0}
						<span class="badge-icon variant-filled-warning absolute -top-0 -right-0 z-10">
								{empresas}
						</span>
					{/if}
			</li>
		</ul>
	</nav>
	<!-- --- -->
</svelte:fragment>
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">Venda coberta de Opções</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- Page Route Content -->
	<slot />
</AppShell>
