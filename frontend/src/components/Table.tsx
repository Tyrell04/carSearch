

export function Table({ children }) {
    return (
        <table>
            {children}
        </table>
    );
}

Table.Row = ({column , ...props }) => {
    return (
                <tbody>
                    <tr>
                        {column.map((column) => (
                            <td>{column}</td>
                        ))}
                    </tr>
                </tbody>
    );
};

Table.Header = ({ columns }) => {
    return (
        <thead>
            <tr>
                {columns.map((column) => (
                    <th>{column}</th>
                ))}
            </tr>
        </thead>
    );
}

Table.Foot = (columns) => {
    return (
        <tfoot>
            <tr>
                {columns.map((column, index) => (
                    <td key={index}>{column}</td>
                ))}
            </tr>
        </tfoot>
    );
}