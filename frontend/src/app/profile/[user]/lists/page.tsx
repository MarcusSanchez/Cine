"use client";

import { useState } from "react";

export default function ProfileLists() {
  const [lists, setLists] = useState([]);

  return (
    <div className="container">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      </div>
    </div>
  );

}