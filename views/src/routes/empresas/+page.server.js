import * as backend from '$lib/server/backend';

/** @satisfies {import('./$types').Actions} */
export const actions = {
	update: async ({ request }) => {
		const data = await request.formData();
		var object = [];
		data.forEach(function(value, key){
			if (!value)
				return
			object.push( {
				id: key,
				tipo: value,
			})
		});
		return {	
			data: await backend.patchEmpresasSemTipo(object)
		};
	}
};
