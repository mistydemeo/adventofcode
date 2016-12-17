(reduce +
  (for [c (slurp "input.txt")]
    (if (= c \() 1 -1)))))
