<script lang="ts">
    import { closeModal } from 'svelte-modals';
  import type { IGolyProps } from '../domain/models/Goly';

    export let title: String;
    export let isOpen: boolean;
    export let goly: IGolyProps;
    export let send: (data: IGolyProps) => Promise<void>;

    let data = { ...goly };

    async function handleSend() {
        await send(data);
        closeModal();
    }

</script>

{#if isOpen}
<div role="dialog" class="modal">
    <div class="contents">
        <h2>{title}</h2>

        <label for="redirect">Redirect To:</label>
        <input
            type="text"
            name="redirect"
            id="redirect"
            bind:value="{data.redirect}">

        <label for="goly">Goly</label>
        <input
            type="text"
            name="goly"
            id="goly"
            disabled="{data.random}"
            class="{ data.random ? "disabled" : ""}"
            bind:value="{data.goly}">

        <label for="isRandom">Random Generated Goly</label>
        <input
            type="checkbox"
            name="isRandom" 
            id="isRandom"
            bind:checked="{data.random}">

        <div class="actions">
            <button class="button__cancel" on:click={closeModal}>Cancel</button>
            <button class="button__primary" on:click={handleSend}>Send</button>
        </div>
    </div>
</div>
{/if}

<style>
    .modal {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        display: flex;
        align-items: center;
        justify-content: center;

        box-shadow: 2px 2px 2px #00000099;

        pointer-events: none;
    }

    .contents {
        min-width: 500px;
        border-radius: 6px;
        padding: 16px;
        background: white;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        pointer-events: auto;
    }

    h2 {
        text-align: center;
        font-size: 24px;
    }

    .actions {
        margin-top: 32px;
        display: flex;
        justify-content: space-between;
        gap: 8px;
    }

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

    label {
        margin-bottom: 4px;
    }

    input {
        padding: 8px 4px;
        margin-bottom: 16px;
    }

    input[type=checkbox] {
        justify-content: flex-start ;
    }

    .disabled {
        background: slategray;
    }
</style>