<script lang="ts">
    import { openModal } from 'svelte-modals';

    import Modal from "./Modal.svelte";

    import { GolyService } from "../services/GolyService";
    import type { IGolyProps } from "../domain/models/Goly";

    export let refetch: () => Promise<void>;

    async function handleCreate(newGoly: IGolyProps) {
        await GolyService.create(newGoly);
        refetch();
    }

    function handleOpenCreateModal() {
        const newGoly:IGolyProps = {
            clicked: 0,
            goly: '',
            redirect: '',
            random: false,
        }

        openModal(Modal, {
            title: "New Goly Link",
            send: handleCreate,
            goly: newGoly,
        });
    }
</script>

<div class="container">
    <button class="button__primary" on:click={handleOpenCreateModal}>New</button>
</div>

<style>

    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: 0.75rem;
        border-radius: 4px;

        cursor: pointer;

        margin-left: 24px;
        width: 200px;
    }

    .button__primary {
        background-color: yellowgreen;
    }

    .button__cancel {
        background-color: red;
    }

    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: #00000077;
    }

</style>