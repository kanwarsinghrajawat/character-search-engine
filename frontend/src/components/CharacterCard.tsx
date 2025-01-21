import React from "react";
import { Character } from "../types/Character";

const CharacterCard: React.FC<{ character: Character }> = ({ character }) => {
  return (
    <div className="bg-[#1E293B] p-6 rounded-xl shadow-lg text-center transform hover:scale-105 transition-transform duration-300 border border-gray-700">
      <h2 className="text-2xl font-bold mb-4 text-white">{character.name}</h2>
      <img
        src={character.image}
        alt={character.name}
        className="rounded-xl w-40 h-40 mx-auto mb-4 border-2 border-[#60A5FA] shadow-lg"
      />
      <p className="text-lg text-gray-300">
        <strong className="text-[#34D399]">Status:</strong> {character.status}
      </p>
      <p className="text-lg text-gray-300">
        <strong className="text-[#34D399]">Species:</strong> {character.species}
      </p>
      <p className="text-lg text-gray-300">
        <strong className="text-[#34D399]">Gender:</strong> {character.gender}
      </p>
      {character.origin?.name && (
        <p className="text-lg text-gray-300">
          <strong className="text-[#34D399]">Origin:</strong>{" "}
          {character.origin.name}
        </p>
      )}
      {character.location?.name && (
        <p className="text-lg text-gray-300">
          <strong className="text-[#34D399]">Location:</strong>{" "}
          {character.location.name}
        </p>
      )}
      <p className="text-lg text-gray-300">
        <strong className="text-[#34D399]">Episodes:</strong>{" "}
        {character.episode?.length || 0}
      </p>
    </div>
  );
};

export default CharacterCard;
