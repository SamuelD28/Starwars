import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";

export type PeoplePropertyTableProps = {
  name: string;
  data: Record<string, any>[];
  columns: string[];
};

export const PeoplePropertyTable = ({
  name,
  data,
  columns,
}: PeoplePropertyTableProps) => {
  return (
    <div className="flex-grow-1">
      <h4>{name}</h4>
      <DataTable value={data}>
        {columns.map((x) => (
          <Column
            headerClassName="py-3 px-0 text-sm"
            className="py-3 px-0 text-sm"
            field={x}
            header={x}
          ></Column>
        ))}
      </DataTable>
    </div>
  );
};
