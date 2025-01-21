import React from "react";
import CharacterCard from "./CharacterCard";
import { Character } from "../types/Character";

const CharacterList: React.FC<{ characters: Character[] }> = ({
  characters,
}) => {
  return (
    <div className="mt-8 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 w-full max-w-6xl">
      {characters.map((character) => (
        <CharacterCard key={character.id} character={character} />
      ))}
    </div>
  );
};

export default CharacterList;
