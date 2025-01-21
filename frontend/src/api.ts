import axios from "axios";
import { Character } from "./types/Character";

// const API_BASE_URL = "http://34.138.200.26:8080";
// const API_BASE_URL = "http://34.138.200.26";
const API_BASE_URL = "http://localhost:8080";

export const fetchCharacterByName = async (
  name: string
): Promise<{ characters: Character[] }> => {
  try {
    const encodedName = encodeURIComponent(name.trim());
    const response = await axios.get(
      `${API_BASE_URL}/character/${encodedName}`
    );

    return response.data?.characters ? response.data : { characters: [] };
  } catch (error) {
    console.error("Error fetching character:", error);
    return { characters: [] };
  }
};
