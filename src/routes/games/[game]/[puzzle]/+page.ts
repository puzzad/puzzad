import type {PageLoad} from './$types';

export const load: PageLoad = async ({params}) => {
    return {
        puzzle: params.puzzle,
        game: params.game
    };
}