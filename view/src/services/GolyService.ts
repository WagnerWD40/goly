import type { IGolyProps } from "../domain/models/Goly";

import config from "../config";

export namespace GolyService {

    export async function fetchAll() {
        const res = await fetch(`${config.BASE_URL}/goly`);
        return await res.json();
    }

    export async function create(newGoly: IGolyProps) {
        await fetch(`${config.BASE_URL}/goly`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(newGoly),
        }).then(res => console.log(res));
    }

    export async function update(newData: IGolyProps, oldGoly: IGolyProps) {
        const updatedGoly: IGolyProps = {
            ...newData,
            id: oldGoly.id,
        }

        await fetch(`${config.BASE_URL}/goly`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(updatedGoly),
        }).then(res => console.log(res));
    }

    export async function destroy(goly: IGolyProps) {
        await fetch(`${config.BASE_URL}/goly/${goly.id}`, {
            method: "DELETE",
        }).then(res => console.log(res));
    }
}