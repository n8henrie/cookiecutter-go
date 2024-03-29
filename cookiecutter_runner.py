import datetime
import os.path

from cookiecutter.main import cookiecutter

curdir = os.path.dirname(os.path.abspath(__file__))

today = datetime.datetime.today()
date = today.date().isoformat()
year = today.year

extra_context = {"release_date": date, "year": year}

cookiecutter(curdir, extra_context=extra_context)
