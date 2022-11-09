import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    return {
        adventure: params.adventure
    };
}