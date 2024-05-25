
export interface TableRowData {
    tfoot?: boolean;
    columns: string[];
}

export interface TableProps {
    data: TableRowData[];
    thead?: string[];
}


export function Table({ data, thead, }: TableProps) {
    return (
        <table>
            <thead>
            <tr>
                {thead?.map((column, index) => (
                    <th key={index}>{column}</th>
                ))}
            </tr>
            </thead>
            {data.map((row, index) => (
                <Row key={index} {...row} />
            ))}
        </table>
    );
}

function Row({ columns, tfoot, }: TableRowData) {
    if (tfoot) {
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
    return (
        <tbody>
        <tr>
            {columns.map((column, index) => (
                <td key={index}>{column}</td>
            ))}
        </tr>
        </tbody>
    );
}