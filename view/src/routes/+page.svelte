<script lang="ts">
    import List from "../components/List.svelte";
    import New from "../components/New.svelte";

    import { onMount } from "svelte";

    import { GolyService } from "../services/GolyService";
    import type { IGolyProps } from "../domain/models/Goly";

    let golies: IGolyProps[] = [];

    async function refetch() {
        golies = await GolyService.fetchAll();
    }

    onMount(async () => {
        refetch();
    })
</script>

<main>
    <h1>Goly -- The Go URL Shortener!</h1>
    <New refetch={refetch} />
    <List refetch={refetch} golies={golies} />
</main>

<style>
    html, body, main {
        font-family: sans-serif;
    }
</style>
