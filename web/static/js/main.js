const revealElements = [...document.querySelectorAll("[data-reveal]")];
const yearNode = document.getElementById("year");

if (yearNode) {
  yearNode.textContent = String(new Date().getFullYear());
}

const revealObserver = new IntersectionObserver(
  (entries) => {
    entries.forEach((entry) => {
      if (!entry.isIntersecting) {
        return;
      }
      entry.target.classList.add("is-visible");
      revealObserver.unobserve(entry.target);
    });
  },
  { threshold: 0.16 }
);

revealElements.forEach((el) => revealObserver.observe(el));

const heroCard = document.querySelector(".hero-card");
if (heroCard && window.matchMedia("(pointer:fine)").matches) {
  heroCard.addEventListener("mousemove", (event) => {
    const rect = heroCard.getBoundingClientRect();
    const x = (event.clientX - rect.left) / rect.width;
    const y = (event.clientY - rect.top) / rect.height;
    const rotateY = (x - 0.5) * 6;
    const rotateX = (0.5 - y) * 5;

    heroCard.style.transform = `perspective(900px) rotateX(${rotateX}deg) rotateY(${rotateY}deg)`;
  });

  heroCard.addEventListener("mouseleave", () => {
    heroCard.style.transform = "perspective(900px) rotateX(0deg) rotateY(0deg)";
  });
}

const heroSlider = document.querySelector("[data-slider]");
if (heroSlider) {
  initHeroSlider(heroSlider);
}

hydrateCompany();

function initHeroSlider(root) {
  const slides = [...root.querySelectorAll("[data-slide]")];
  const dots = [...root.querySelectorAll("[data-slider-dot]")];
  const prevButton = root.querySelector("[data-slider-prev]");
  const nextButton = root.querySelector("[data-slider-next]");
  let activeIndex = 0;
  let touchStartX = 0;

  const setActiveSlide = (index) => {
    const boundedIndex = (index + slides.length) % slides.length;
    activeIndex = boundedIndex;

    slides.forEach((slide, slideIndex) => {
      slide.classList.toggle("is-active", slideIndex === boundedIndex);
    });

    dots.forEach((dot, dotIndex) => {
      dot.classList.toggle("is-active", dotIndex === boundedIndex);
      dot.setAttribute("aria-current", dotIndex === boundedIndex ? "true" : "false");
    });
  };

  prevButton?.addEventListener("click", () => setActiveSlide(activeIndex - 1));
  nextButton?.addEventListener("click", () => setActiveSlide(activeIndex + 1));

  dots.forEach((dot, index) => {
    dot.addEventListener("click", () => setActiveSlide(index));
  });

  root.addEventListener("touchstart", (event) => {
    touchStartX = event.changedTouches[0]?.clientX ?? 0;
  }, { passive: true });

  root.addEventListener("touchend", (event) => {
    const touchEndX = event.changedTouches[0]?.clientX ?? 0;
    const distance = touchEndX - touchStartX;

    if (Math.abs(distance) < 40) {
      return;
    }

    setActiveSlide(distance > 0 ? activeIndex - 1 : activeIndex + 1);
  }, { passive: true });

  setActiveSlide(0);
}

async function hydrateCompany() {
  try {
    const response = await fetch("/api/company", { headers: { Accept: "application/json" } });
    if (!response.ok) {
      throw new Error("company request failed");
    }

    const company = await response.json();

    setText("legalName", company.legal_name);
    setText("inn", company.inn);
    setText("kpp", company.kpp);
    setText("legalAddress", company.legal_address);
    setText("postalAddress", company.postal_address);
    setText("settlementAccount", company.settlement_account);
    setText("bankName", company.bank_name);
    setText("correspondentAccount", company.correspondent_account);
    setText("bik", company.bik);
    setText("director", company.director);
    setText("address", company.legal_address);
    setText("registryPhone", company.phone);
    setText("registryEmail", company.email);

    if (company.phone) {
      const phone = document.getElementById("phone");
      if (phone) {
        phone.textContent = company.phone;
        phone.href = `tel:${company.phone.replace(/[^\d+]/g, "")}`;
      }
    }

    if (company.email) {
      const mail = document.getElementById("email");
      if (mail) {
        mail.textContent = company.email;
        mail.href = `mailto:${company.email}`;
      }
    }

  } catch (error) {
    // Реквизиты остаются в статическом HTML, если API временно недоступен.
  }
}

function setText(id, value) {
  const node = document.getElementById(id);
  if (node && value) {
    node.textContent = value;
  }
}
