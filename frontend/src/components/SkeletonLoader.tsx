import React from "react";
import SkeletonCard from "./SkeletolCard";

const SkeletonLoader: React.FC = () => {
  return (
    <div className="mt-8 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 w-full max-w-6xl">
      {Array(6)
        .fill(0)
        .map((_, index) => (
          <SkeletonCard key={index} />
        ))}
    </div>
  );
};

export default SkeletonLoader;
