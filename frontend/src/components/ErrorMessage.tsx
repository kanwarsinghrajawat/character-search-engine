import React from "react";

const ErrorMessage: React.FC<{ message: string }> = ({ message }) => {
  return (
    <div className="mt-6 flex flex-col items-center">
      <img
        src="https://cdn-icons-png.flaticon.com/512/2748/2748558.png"
        alt="Not Found"
        className="w-32 h-32 opacity-80 animate-bounce"
      />
      <p className="text-red-500 text-lg font-semibold mt-4">{message}</p>
    </div>
  );
};

export default ErrorMessage;
