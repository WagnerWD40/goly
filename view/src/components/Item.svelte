<script lang="ts">
    import type { IGolyProps } from "../domain/models/Goly";
    import { Modals, closeModal, openModal } from 'svelte-modals';

    import { GolyService } from "../services/GolyService";
    
    import Card from "./Card.svelte";
    import Modal from "./Modal.svelte";
    
    import config from "../config";

    export let goly: IGolyProps;
    export let refetch: () => Promise<void>;

    async function handleUpdate(data: IGolyProps) {
        await GolyService.update(data, goly);
        refetch();
    }

    async function handleDelete() {
        if(confirm(`Are yousure you wish to delete this Goly link (${goly.goly})`)) {
            await GolyService.destroy(goly);
            refetch();
        }
    }

    function handleOpenUpdateModal() {
        openModal(Modal, {
            title: "Update Goly Link",
            send: handleUpdate,
            goly: goly,
        });
    }

    const link = `${config.BASE_URL}/r/${goly.goly}`;
</script>

<Card>
    <p>Goly: <a href="{link}" target="_blank">{link}</a></p>
    <p>Redirect: {goly.redirect}</p>
    <p>Clicked: {goly.clicked} </p>
    <button class="button__primary" on:click={handleOpenUpdateModal}>Update</button>
    <button class="button__cancel" on:click={handleDelete}>Delete</button>
</Card>
<Modals>
    <div
        slot="backdrop"
        class="backdrop"
        on:click={closeModal} />
</Modals>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: 0.75rem;
        border-radius: 4px;

        cursor: pointer;
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