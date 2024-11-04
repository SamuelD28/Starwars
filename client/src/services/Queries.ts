import { gql, useLazyQuery } from "@apollo/client";
import { People } from "../models/People";

const getPeoplesByName = gql`
  query GetPeoplesByName($search: String!) {
    peoples(search: $search) {
        name
        id
        gender
    }
  }
`;

export const useGetPeoplesByName = () => useLazyQuery<{ peoples: People[] }>(getPeoplesByName);

const getPeopleById = gql`
  query GetPeopleById($id: Int!) {
    people(id: $id) {
        name
        id
        gender
        vehicles {
          name
          model
          vehicleClass
          manufacturer
        }
        films {
          title
          episodeId
          releaseDate
        }
    }
  }
`;


export const useGetPeopleById = () => useLazyQuery<{ people: People }>(getPeopleById)