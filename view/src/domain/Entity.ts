type BasicEntity = {
    id: number;
}

abstract class Entity<T extends BasicEntity> {
    private readonly props: T;

    get id(): number {
        return this.props.id;
    }

    get key(): string {
        return `${this.props.id}`;
    }

    get marshall(): T {
        return this.props;
    }

    protected constructor(props: T) {
        this.props = props;
    }
}

export default Entity;