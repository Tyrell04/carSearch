export enum InputType {
    TEXT = 'text',
    PASSWORD = 'password',
    EMAIL = 'email',
    NUMBER = 'number'
}

export function Form({ children, onSubmit }) {
    return (
        <form onSubmit={onSubmit}>
            {children}
        </form>
    );
}

Form.Input = ({ type = InputType.TEXT, name, ...props }) => {
    return (
        <input type={type} name={name} {...props} />
    );
};

Form.Submit = ({ children }) => {
    return (
        <button type="submit">{children}</button>
    );
}
