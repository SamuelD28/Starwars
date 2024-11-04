import { PrimeReactProvider } from "primereact/api";
import { useEffect, useState } from "react";
import {
  AutoComplete,
  AutoCompleteCompleteEvent,
  AutoCompleteSelectEvent,
  AutoCompleteUnselectEvent,
} from "primereact/autocomplete";
import { People } from "./models/People";
import { useGetPeopleById, useGetPeoplesByName } from "./services/Queries";
import { PeopleCard } from "./components/PeopleCard";

function App() {
  const [search, setSearch] = useState("");
  const [selected, setSelected] = useState<People[]>([]);
  const getPeoplesByNameQuery = useGetPeoplesByName();
  const [suggestions, setSuggestions] = useState<People[]>([]);
  const getPeopleByIdQuery = useGetPeopleById();

  useEffect(() => {
    const body = document.querySelector("body");
    const banner = document.querySelector<HTMLElement>("#banner");
    body?.addEventListener("mousemove", function (e) {
      var amountMovedX = (e.pageX * -1) / 25;
      var amountMovedY = (e.pageY * -1) / 25;
      if (banner) {
        banner.style.backgroundPosition =
          amountMovedX + "px " + amountMovedY + "px";
      }
    });
  });

  const onSelect = async (event: AutoCompleteSelectEvent) => {
    const selectedPeople = event.value as People;

    if (selected.findIndex((x) => x.id == selectedPeople.id) !== -1) {
      return;
    }

    const result = await getPeopleByIdQuery[0]({
      variables: { id: selectedPeople.id },
    });
    if (!result.error && result.data) {
      setSelected([result.data.people, ...selected]);
    }
  };

  const onUnSelect = async (event: AutoCompleteUnselectEvent) => {
    const selectedPeople = event.value as People;
    setSelected([...selected.filter((x) => x.id !== selectedPeople.id)]);
  };

  const onSearch = async (event: AutoCompleteCompleteEvent) => {
    const result = await getPeoplesByNameQuery[0]({
      variables: { search: event.query },
    });

    const alreadySelectedIds = selected.map((x) => x.id);
    if (result.data?.peoples) {
      setSuggestions(
        result.data.peoples.filter((x) => !alreadySelectedIds.includes(x.id))
      );
    }
  };
  return (
    <PrimeReactProvider>
      <div id="banner"></div>
      <section className="flex">
        <div
          style={{ width: "40%" }}
          className="flex flex-column align-items-center p-5 mt-7"
        >
          <div>
            <img width="275px" height="275px" src="logo.png" />
          </div>
          <AutoComplete
            className="w-full"
            id="username"
            pt={{
              container: { className: "flex-grow-1 p-3" },
              token: { className: "p-3" },
            }}
            placeholder="Search character..."
            showEmptyMessage
            delay={500}
            multiple
            value={search}
            onChange={(e) => setSearch(e.value)}
            emptyMessage="No results found"
            field="name"
            suggestions={suggestions}
            completeMethod={onSearch}
            onSelect={onSelect}
            onUnselect={onUnSelect}
          />
        </div>
        <div
          style={{ width: "60%", height: "95vh" }}
          className="flex flex-column pt-5 pr-5 overflow-y-scroll"
        >
          {selected.map(x => (
            <PeopleCard people={x} />
          ))}
        </div>
      </section>
    </PrimeReactProvider>
  );
}

export default App;
