<script>
	import { enhance, applyAction } from '$app/forms';
	import empresaStore from '../../Store.js';
	import { getToastStore } from '@skeletonlabs/skeleton';

	const toastStore = getToastStore();

	let empresas = []
	
	empresaStore.subscribe((dataSub) => {
		empresas = dataSub;
	});

</script>

<div class="container mx-auto p-8 space-y-8 ">
	<div class="table-container">
		{#if empresas.length > 0}
		<form method="POST" action="?/update"
			use:enhance={({ formElement, formData, action, cancel, submitter }) => {
				return async ({ result, update }) => {
					if (result.type === 'error') {
						await applyAction(result);
					}
					empresaStore.update((currentData) => {
						let _empresas = result?.data?.data
						if (_empresas.length === empresas.length) {
							toastStore.trigger({
								message: 'Não foi possivel salvar. Tente novamente.',
								background: 'variant-filled-error',
							})
						} else {
							toastStore.trigger({
								message: 'Tipo da Ação salva com sucesso.',
								background: 'variant-filled-success',
							})
						}
						return _empresas;
					});
				};
			}}
			>
			<table class="table table-hover table-comfortable">
				<thead>
					<tr>
						<th >Ticker</th>
						<th >Razão Social</th>
						<th >Tipo</th>
					</tr>
				</thead>
				<tbody>
					{#each empresas as empresa}
						<tr>
							<td>{empresa.id}</td>
							<td>{empresa.razSocEmi}</td>
							<td>
								<input name="{empresa.id}" width="10px" class="input w-20" type="text" bind:value={empresa.tipo} />
							</td>
						</tr>
					{/each}
				</tbody>
				<tfoot>
					<tr>
						<td>
						</td>
						<td></td>
						<td >
							<button class="btn variant-ghost-primary">
								<span><i class="fa-solid fa-floppy-disk"></i></span>
								<span>Salvar</span>
							</button>
						</td>
					</tr>
				</tfoot>
			</table>
		</form>
		{:else}
		<div class="card p-4">
			Nenhuma inconsistencia encontrada.
		</div>
		{/if}
	</div>

</div>
