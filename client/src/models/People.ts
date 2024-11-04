import { Film } from "./Film"
import { Vehicle } from "./Vehicle"

export type People = {
  name: string
  id: number
  gender: string
  films: Film[]
  vehicles: Vehicle[]
}