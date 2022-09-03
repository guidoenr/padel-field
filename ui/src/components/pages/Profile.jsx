import React from "react";
import { motion } from "framer-motion";

const Profile = () => {
  const containerVariants = {
    hidden: {
      opacity: 0,
      x: "100vw",
    },
    visible: {
      opacity: 1,
      x: 0,
      transition: {
        type: "spring",
        bounce: 0.25,
      },
    },
    exit: {
      x: "-100vw",
      transition: {
        type: "spring",
        bounce: 0.25,
      },
    },
  };

  return (
    <section className="profile h-[80vh] py-8 w-full">
      <motion.div
        variants={containerVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
        className="profile-container container mx-auto px-1 flex flex-col"
      >
        <div className="profile-header relative py-4">
          <h3 className="text-2xl lg:text-3xl font-medium">Mi cuenta:</h3>
          <div className="absolute bottom-1 w-[3rem] h-[.25rem] bg-accent"></div>
          <div className="absolute bottom-[-.25rem] w-[1.5rem] h-[.25rem] bg-accent"></div>
        </div>
      </motion.div>
    </section>
  );
};

export default Profile;
