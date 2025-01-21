import React from "react";
import { FaSearch } from "react-icons/fa";

interface SearchBarProps {
  name: string;
  setName: (name: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ name, setName }) => {
  return (
    <div className="flex items-center space-x-3 w-full max-w-xl relative">
      <input
        type="text"
        placeholder="Search character..."
        value={name}
        onChange={(e) => setName(e.target.value)}
        className="w-full p-4 pl-12 text-lg rounded-full text-black border border-gray-500 bg-gray-200 focus:outline-none focus:ring-4 focus:ring-[#60A5FA] transition-all duration-300 shadow-md"
      />
      <FaSearch className="absolute left-4 text-gray-500 text-lg" />
    </div>
  );
};

export default SearchBar;
