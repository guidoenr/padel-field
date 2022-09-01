import React from "react";

const Profile = () => {
  return (
    <section className="profile h-[80vh] py-8 w-full">
      <div className="profile-container container mx-auto px-1 flex flex-col">
        <div className="profile-header relative py-4">
          <h3 className="text-2xl lg:text-3xl font-medium">Mi cuenta:</h3>
          <div className="absolute bottom-1 w-[3rem] h-[.25rem] bg-accent"></div>
          <div className="absolute bottom-[-.25rem] w-[1.5rem] h-[.25rem] bg-accent"></div>
        </div>
      </div>
    </section>
  );
};

export default Profile;
