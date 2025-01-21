import React, { useState, useEffect } from "react";
import { Character } from "./types/Character";
import CharacterList from "./components/CharacterList";
import ErrorMessage from "./components/ErrorMessage";
import SkeletonLoader from "./components/SkeletonLoader";
import { fetchCharacterByName } from "./api";
import SearchBar from "./components/Searchbar";

const SearchCharacter: React.FC = () => {
  const [name, setName] = useState<string>("");
  const [characters, setCharacters] = useState<Character[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [debouncedName, setDebouncedName] = useState<string>(name);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedName(name);
    }, 300);

    return () => clearTimeout(handler);
  }, [name]);

  useEffect(() => {
    if (debouncedName.trim() === "") {
      setCharacters([]);
      setError("Please enter a character name.");
      return;
    }

    const fetchCharacters = async () => {
      setLoading(true);
      setError(null);

      try {
        const results = await fetchCharacterByName(debouncedName);
        const charactersArray = results?.characters || [];

        if (charactersArray.length > 0) {
          setCharacters(charactersArray);
        } else {
          setCharacters([]);
          setError("No characters found.");
        }
      } catch (err) {
        console.error("Error fetching character:", err);
        setError("Failed to fetch character.");
      } finally {
        setLoading(false);
      }
    };

    fetchCharacters();
  }, [debouncedName]);

  return (
    <div className="min-h-screen bg-[#0F172A] text-white flex flex-col items-center p-6">
      <h1 className="text-4xl font-bold mb-6 text-center text-[#60A5FA] flex items-center gap-2">
        🔍 Rick and Morty Character Search
      </h1>

      <SearchBar name={name} setName={setName} />

      {loading && <SkeletonLoader />}

      {error && !loading && <ErrorMessage message={error} />}

      {!error && !loading && <CharacterList characters={characters} />}
    </div>
  );
};

export default SearchCharacter;
