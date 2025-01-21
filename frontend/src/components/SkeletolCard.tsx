import React from "react";

const SkeletonCard: React.FC = () => {
  return (
    <div className="bg-[#1E293B] p-6 rounded-xl shadow-lg animate-pulse border border-gray-700">
      <div className="h-6 bg-gray-700 rounded w-3/4 mb-4"></div>
      <div className="w-40 h-40 mx-auto mb-4 bg-gray-700 rounded-lg"></div>
      <div className="h-4 bg-gray-700 rounded w-1/2 mb-2"></div>
      <div className="h-4 bg-gray-700 rounded w-3/4 mb-2"></div>
      <div className="h-4 bg-gray-700 rounded w-1/3 mb-2"></div>
      <div className="h-4 bg-gray-700 rounded w-2/3"></div>
    </div>
  );
};

export default SkeletonCard;
