// Get the modal
const modal = document.getElementById("cityModal");
const closeBtn = document.querySelector(".close");

// Get modal elements
const cityName = document.getElementById("cityName");
const cityDescription = document.getElementById("cityDescription");

// City data
const cities = {
  Bandung: "Bandung adalah kota yang dikenal dengan cuaca sejuk dan kuliner khas seperti batagor dan siomay.",
  Jakarta: "Jakarta, ibukota Indonesia, memiliki beragam kuliner seperti kerak telor dan soto Betawi.",
  Jogja: "Yogyakarta terkenal dengan kuliner tradisionalnya seperti gudeg dan bakpia.",
  Semarang : "Kota Semarang adalah ibu kota provinsi Jawa Tengah, Indonesia. Kota ini adalah kota metropolitan terbesar kelima di Indonesia setelah Jakarta, Surabaya, Medan, dan Bandung.",
  Surabaya : "Kota Surabaya adalah ibu kota Provinsi Jawa Timur yang menjadi pusat pemerintahan dan perekonomian sekaligus kota terbesar di provinsi tersebut. Surabaya juga merupakan sebuah kota yang terletak di Provinsi Jawa Timur, Indonesia.",
  Bali : "Bali adalah sebuah provinsi di Indonesia yang terletak pada bagian barat Kepulauan Nusa Tenggara dan beribu kota di Kota Denpasar. Pulau Bali, yang merupakan pulau terbesar di Provinsi Bali, memiliki beberapa julukan, di antaranya Pulau Dewata dan Pulau Seribu Pura.",
  Malang : "Kota Malang adalah sebuah kota yang terletak di Provinsi Jawa Timur, Indonesia, Kota ini merupakan kota terbesar kedua di Jawa Timur setelah Surabaya, dan kota terbesar ke-12 di Indonesia. Kota ini didirikan pada masa Pemerintahan Belanda pada 1 April 1914 dengan E.K Broeveldt sebagai wali kota pertama.",
  Medan : "Medan, ibu kota Sumatera Utara, Indonesia, adalah pusat ekonomi utama dan tempat perpaduan budaya, yang memadukan pengaruh Melayu, Batak, Tiongkok, dan India.",
};

// Add event listener for "See More" buttons
document.querySelectorAll(".seemore_bt a").forEach(btn => {
  btn.addEventListener("click", function(event) {
    event.preventDefault();
    
    // Get the city name
    const city = this.parentElement.parentElement.querySelector("h4").textContent;
    
    // Set modal content
    cityName.textContent = city;
    cityDescription.textContent = cities[city] || "Deskripsi tidak tersedia.";

    // Show the modal
    modal.style.display = "flex";
  });
});

// Close modal when 'X' is clicked
closeBtn.addEventListener("click", () => {
  modal.style.display = "none";
});

// Close modal if user clicks outside the modal content
window.addEventListener("click", (event) => {
  if (event.target === modal) {
    modal.style.display = "none";
  }
  
});
