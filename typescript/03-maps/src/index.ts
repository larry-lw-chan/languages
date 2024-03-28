import { User } from "./User";
import { Company } from "./Company";
import { CustomMap } from "./CustomMap";

// Create user and company instance
const user = new User();
const company = new Company();

// Create map instance
const mapDiv = document.getElementById("map");

if (mapDiv !== null) {
  const customMap = new CustomMap(mapDiv);
  customMap.addMarker(user);
  customMap.addMarker(company);
}
