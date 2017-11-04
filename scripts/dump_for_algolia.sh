echo "id,slug,brand,model,category" > lowside-io-algolia.csv
docker exec -i core_database_1 su postgres -c 'psql --dbname=lowside-dev -P format=unaligned -P tuples_only -P fieldsep=\, -c "SELECT id, slug, brand, model, category FROM public.\"motorcycles_models\""' >> lowside-io-algolia.csv
