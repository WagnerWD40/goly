<script lang="ts">
    import Item from "./Item.svelte";

    import type { IGolyProps } from "../domain/models/Goly";

    export let golies: IGolyProps[];
    export let refetch: () => Promise<void>

</script>

{#each golies as goly}
    <Item goly={goly} refetch={refetch} />
{/each}