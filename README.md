# pinda

IBS3 experimentele backend

# Init

Om te beginnen met ontwikkelen moet je eerst twee dingen doen

1. `task init`
2. Daarna, moet je in het bestand /ibs3/prisma/schema.prisma de volgende lijn vervangen:

```
generator client {
  provider = "prisma-client-js"
}
```

naar

```
generator client {
  provider = "go run github.com/steebchen/prisma-client-go"
  output = "../../db"
}
```

_niet committen lmao_
