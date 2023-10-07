import * as backend from '$lib/server/backend';

/** @satisfies {import('./$types').PageServerLoad} */
export const load = () => {
	return {
		ativos: backend.fetchAtivos()
	};
};