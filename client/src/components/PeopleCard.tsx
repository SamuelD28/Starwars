import { Card } from "primereact/card";
import { People } from "../models/People";
import { PeoplePropertyTable } from "./PeoplePropertyTable";

export type PeopleCardProps = {
  people: People;
};

export const PeopleCard = ({
  people: { name, vehicles, films },
}: PeopleCardProps) => {
  return (
    <Card className="w-full mb-5">
      <h3 className="m-0 text-primary">{name}</h3>
      <div className="flex">
        <PeoplePropertyTable
          name="Vehicles"
          data={vehicles}
          columns={["name", "model"]}
        />
        <PeoplePropertyTable
          name="Films"
          data={films}
          columns={["title", "episodeId"]}
        />
      </div>
    </Card>
  );
};
